package service

import (
	"binginx.com/brush/config"
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/internal/logs"
	"binginx.com/brush/internal/service"
	"binginx.com/brush/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func UserInfo(user *model.UserInfo, params *config.Params) (err error, jicUser *model.JicUser) {
	request, err := New(user, _userInfo_url, _method_get, params)
	if err != nil {
		return err, jicUser
	}
	resp, err1 := clients.Excute(request)
	if err1 != nil {
		logs.Logger.Errorf(err1.Error())
		return err1, jicUser
	}
	defer resp.Body.Close()
	out, err2 := ioutil.ReadAll(resp.Body)
	err2 = json.Unmarshal(out, &jicUser)
	if err2 != nil {
		logs.Logger.Errorf(err2.Error())
		return err2, jicUser
	}
	if jicUser.Code != http.StatusOK {
		return errors.New(jicUser.Msg), nil
	}
	getUser, err := service.GetUserByPhone(jicUser.JicUserInfo.Phone)
	if getUser == nil {
		service.CreateUser(&model.User{
			Name:       jicUser.JicUserInfo.Name,
			Token:      user.Token,
			Phone:      jicUser.JicUserInfo.Phone,
			EmployeeID: jicUser.JicUserInfo.EmployeeID,
		})
	} else {
		service.UpdateUser(getUser)
	}
	return err, jicUser
}

func Score(user *model.UserInfo, params *config.Params) (err error, score string) {
	request, err := New(user, _token_url, _method_get, params)
	if err != nil {
		return err, score
	}
	resp, err1 := clients.Excute(request)
	if err1 != nil {
		logs.Logger.Errorf(err1.Error())
		return err1, score
	}
	defer resp.Body.Close()
	out, _ := ioutil.ReadAll(resp.Body)
	var scoreInfo model.Resp
	err2 := json.Unmarshal(out, &scoreInfo)
	if err2 != nil {
		logs.Logger.Errorf(err2.Error())
		return err2, score
	}
	if scoreInfo.Code != http.StatusOK {
		return errors.New(scoreInfo.Msg), score
	}
	//getUser, err3 := service.GetUserByToken(user.Token)
	//if err3 != nil {
	//	return err3, score
	//}
	//if getUser != nil {
	//service.UpdateUserScore(getUser.Phone, scoreInfo.Data)
	//}
	return err, scoreInfo.Data
}

func News(user *model.UserInfo, params *config.Params) (err error, ecosysNewsId []model.Content) {
	request, err := New(user, _news_url, _method_get, params)
	if err != nil {
		return err, nil
	}
	resp, err1 := clients.Excute(request)
	if err1 != nil {
		logs.Logger.Errorf(err1.Error())
		return err1, nil
	}
	defer resp.Body.Close()
	out, _ := ioutil.ReadAll(resp.Body)
	var ecosysNews model.Contents
	err2 := json.Unmarshal(out, &ecosysNews)
	if err2 != nil {
		logs.Logger.Errorf(err2.Error())
		return err2, nil
	}
	if ecosysNews.Code != http.StatusOK {
		return errors.New(ecosysNews.Msg), nil
	}
	if ecosysNews.Data != nil && len(ecosysNews.Data) > 0 {
		return nil, ecosysNews.Data
	}
	return err, nil
}

func View(userInfo *model.UserInfo, params *config.Params) (err error) {
	request, err := New(userInfo, _view_url, _method_get, params)
	if err != nil {
		return err
	}
	resp, err1 := clients.Excute(request)
	if err1 != nil {
		log.Println(err1)
		return err1
	}
	defer resp.Body.Close()
	out, err2 := ioutil.ReadAll(resp.Body)
	var viewResult model.ViewContent
	err2 = json.Unmarshal(out, &viewResult)
	if err2 != nil {
		logs.Logger.Errorf(err2.Error())
		return err2
	}
	if viewResult.Code != http.StatusOK {
		return errors.New(viewResult.Msg)
	}
	logs.Logger.Infof("View response status:[%v],contentId:[%v]", resp.Status, params.Params["contentId"])
	return
}

func DoBehavior(userInfo *model.UserInfo, params *config.Params) (err error) {
	request, err := New(userInfo, _behavior_url, _method_post, params)
	if err != nil {
		return err
	}
	resp, err1 := clients.Excute(request)
	if err1 != nil {
		logs.Logger.Errorf(err1.Error())
		return err1
	}
	defer resp.Body.Close()
	out, err2 := ioutil.ReadAll(resp.Body)
	logs.Logger.Infof(string(out))
	var behaviorResult model.Resp
	err2 = json.Unmarshal(out, &behaviorResult)
	if err2 != nil {
		logs.Logger.Errorf(err2.Error())
		return err2
	}
	if behaviorResult.Code != http.StatusOK {
		return errors.New(behaviorResult.Msg)
	}
	logs.Logger.Infof("DoBehavior response: status [%v],contentId:[%v],flag:[%v]", resp.Status, params.Params["resourceId"], params.Params["userBehavior"])
	return
}
