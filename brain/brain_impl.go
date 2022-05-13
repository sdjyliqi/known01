package brain

import (
	"github.com/gansidui/ahocorasick"
	"github.com/golang/glog"
	"known01/model"
	"known01/utils"
	"log"
	"strings"
)

//init() ...初始化鉴别引擎数据
func (c *Center) init() error {
	//load templates about bank
	err := c.InitTemplatesItemsFromDB()
	if err != nil {
		glog.Errorf("Call InitTemplatesItemsFromDB failed,err:%+v", err)
		return err
	}
	//load cut-words from mysql
	err = c.InitCutWordsFromDB()
	if err != nil {
		glog.Errorf("Call InitCutWordsFromDB failed,err:%+v", err)
		return err
	}

	err = c.InitReferencesItemsFromDB()
	if err != nil {
		glog.Errorf("Call InitReferencesItemsFromDB failed,err:%+v", err)
		return err
	}

	return nil
}

//InitCutWordsFromDB ...初始化辅助词，数据来源mysql
func (c *Center) InitCutWordsFromDB() error {
	items, err := model.SpecialSymbolModel.GetItems()
	if err != nil {
		return err
	}
	words := make([]string, len(items))
	for k, v := range items {
		words[k] = v.SpecialCharacter
	}
	//构建副助词匹配自动机
	c.cutWords = words
	acCutWord := ahocorasick.NewMatcher()
	acCutWord.Build(words)
	c.acCutWords = acCutWord
	return nil
}

//getReferencesItemsFromDB ...初始化鉴别基准数据
func (c *Center) InitReferencesItemsFromDB() error {
	//初始化customerPhoneDic
	phoneNumsDic := map[string][]*model.DsisEnterpriseBasic{}
	items, err := model.EnterpriseBasicModel.GetItems(utils.GetMysqlClient())
	if err != nil {
		return err
	}
	if len(items) == 0 {
		log.Fatal("The count of items from table reference is zero,please check the reference table in mysql.")
	}
	c.referencesItems = items //尽可能的复用此数据，交付给鉴别引擎
	//定义全量电话号码
	var allPhoneIDs []string
	for _, v := range items {
		//初始化电话号码到详情的映射表
		if len(v.Phone) > 0 {
			phone := strings.ReplaceAll(v.Phone, "-", "")
			_, ok := phoneNumsDic[phone]
			if ok {
				phoneNumsDic[phone] = append(phoneNumsDic[phone], v)
			} else {
				phoneNumsDic[phone] = []*model.DsisEnterpriseBasic{v}
			}
			allPhoneIDs = append(allPhoneIDs, phone)
		}
		if len(v.ManualPhone) > 0 {
			phonesLine := strings.ReplaceAll(v.ManualPhone, "，", ",")
			phonesLine = strings.ReplaceAll(v.ManualPhone, "-", "")
			phoneIDs := strings.Split(phonesLine, ",")
			allPhoneIDs = append(allPhoneIDs, phoneIDs...)
			for _, phone := range phoneIDs {
				_, ok := phoneNumsDic[phone]
				if ok {
					phoneNumsDic[phone] = append(phoneNumsDic[phone], v)
				} else {
					phoneNumsDic[phone] = []*model.DsisEnterpriseBasic{v}
				}
			}
		}
		//扩充分类关键词
		indexWordDic[utils.EngineBank] = append(indexWordDic[utils.EngineBank], v.Name)
		if len(v.AliasNames) > 0 {
			names := strings.Split(v.AliasNames, ",")
			indexWordDic[utils.EngineBank] = append(indexWordDic[utils.EngineBank], names...)
		}
		//把别名和昵称初始化到分类词表中
	}
	c.customerPhoneDic = phoneNumsDic
	//create customer phone numbers for ac
	acPhoneIDs := ahocorasick.NewMatcher()
	c.customerPhones = allPhoneIDs
	acPhoneIDs.Build(allPhoneIDs)
	c.acCustomerPhoneMatch = acPhoneIDs

	//创建兜底基于关键词的匹配自动机
	var indexWords []string
	for k, v := range indexWordDic {
		indexWords = append(indexWords, v...)
		for _, vv := range v {
			c.indexWordsDic[vv] = k
		}
	}
	acIndexWord := ahocorasick.NewMatcher()
	acIndexWord.Build(indexWords)
	c.indexWords = indexWords
	c.acIndexWords = acIndexWord
	return nil
}

//InitTemplatesItemsFromDB ...初始化短信模板相关的内容
func (c *Center) InitTemplatesItemsFromDB() error {
	templateDic := map[string]utils.EngineType{}
	items, err := model.MessageTMPModel.GetItems()
	if err != nil {
		return err
	}
	if len(items) == 0 {
		glog.Fatal("The count of items from table templates is zero,please check the templates table in mysql.")
	}
	//首次上线只上线金融场景
	for _, v := range items {
		if v.Status == 1 {
			templateDic[v.Detail] = utils.EngineBank
		}
	}
	c.messageTemplates = templateDic
	c.messageTemplatesItems = items
	return nil
}

//cutSpecialMessage ...为了统一化处理，剔除‘ ’，‘-’等符合
func (c *Center) cutSpecialMessage(msg string) string {
	msg = strings.ReplaceAll(msg, " ", "")
	msg = strings.ReplaceAll(msg, "-", "")
	return msg
}

//amendMessage ...模板匹配前，需要提出辅助词
func (c *Center) amendMessage(msg string) string {
	//删除字母或者数字
	var rWords []rune
	for _, v := range []rune(msg) {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			continue
		}
		rWords = append(rWords, v)
	}
	amendMessage := string(rWords)
	//开启副助词ac自动机匹配，然后删除
	matchIndex := c.acCutWords.Match(string(amendMessage))
	for _, v := range matchIndex {
		amendMessage = strings.ReplaceAll(amendMessage, c.cutWords[v.Index], "")
	}
	return amendMessage
}

//acFindPhoneNum ...通过AC自动机寻找客服电话
func (c *Center) acFindPhoneID(msg string) (string, bool) {
	matchIndex := c.acCustomerPhoneMatch.Match(msg)
	if len(matchIndex) > 0 {
		return c.customerPhones[matchIndex[0].Index], true
	}
	return "", false
}

//acFindPhoneNum ...根据寻找的客服电话，查找匹配的引擎名称,如果客服电话匹配多个基准单位，临时选择第一个单位对应的鉴别引擎。
func (c *Center) getEngineByPhoneID(phone string) (utils.EngineType, bool) {
	v, ok := c.customerPhoneDic[phone]
	if ok && len(v) > 0 {
		return utils.EngineBank, true
	}
	return utils.EngineUnknown, false
}

//acFindPhoneNum ...寻找关键字
func (c *Center) acFindIndexWord(msg string) (string, bool) {
	matchIndex := c.acIndexWords.Match(msg)
	if len(matchIndex) > 0 {
		return c.indexWords[matchIndex[0].Index], true
	}
	return "", false
}

//acFindPhoneNum ...根据寻找的客服电话，查找匹配的引擎名称
func (c *Center) getEngineByIndexWord(index string) (utils.EngineType, bool) {
	v, ok := c.indexWordsDic[index]
	return v, ok
}

//bankEngineRate ...计算银行类信息匹配度
func (c *Center) matchEngineRate(msg string) (utils.EngineType, float64) {
	matchRate := 0.0
	engineName := utils.EngineUnknown
	for _, v := range c.messageTemplatesItems {
		simValue := utils.SimHashTool.Hash(msg)
		rate := utils.SimHashTool.Similarity(simValue, v.SimHash)
		if rate > matchRate {
			matchRate = rate
			engineName = utils.EngineBank
		}
	}
	return engineName, matchRate
}

//GetEngineName ... 根据提交的信息，判断最符合那个鉴别引擎
func (c *Center) GetEngineName(msg string) (utils.EngineType, string) {
	minMatchLevel := 0.6
	//第一步，判断是否有官方电话号码,如果找到，返回类型和电话即可。
	msg = c.cutSpecialMessage(msg)
	phoneID, ok := c.acFindPhoneID(msg)
	if ok {
		engineName, ok := c.getEngineByPhoneID(phoneID)
		if ok {
			return engineName, phoneID
		}
	}
	//第二步，寻找关键字
	indexWord, ok := c.acFindIndexWord(msg)
	if ok {
		engineName, ok := c.getEngineByIndexWord(indexWord)
		if ok {
			return engineName, ""
		}
	}
	//第三部，修正短信数据，剔除副助词，英文字母或者数字
	amendMessage := c.amendMessage(msg)
	//第四步  顺序匹配模板，选择匹配最高分
	engineName, score := c.matchEngineRate(amendMessage)
	if score > minMatchLevel {
		return engineName, ""
	}
	return utils.EngineReward, ""
}

//JudgeMessage ... 鉴别短信的入口
func (c *Center) JudgeMessage(msg, sender string) (int, *model.DsisEnterpriseBasic) {
	msg = c.cutSpecialMessage(msg)
	engineName, phoneID := c.GetEngineName(msg)
	switch engineName {
	case utils.EngineBank:
		return c.bank.JudgeMessage(msg, phoneID, sender)
	}
	return utils.OutsideKnown, nil
}
