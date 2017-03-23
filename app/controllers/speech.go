package controllers

import (
	"encoding/json"
	"errors"
	"openss/api/app/models"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type SpeechController struct {
	*revel.Controller
}

func (c SpeechController) Index() revel.Result {
	var (
		speechs []models.Speech
		err     error
	)
	speechs, err = models.GetSpeeches()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 200
	return c.RenderJson(speechs)
}

func (c SpeechController) Show(id string) revel.Result {
	var (
		speech   models.Speech
		err      error
		speechID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid speech id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	speechID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid speech id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	speech, err = models.GetSpeech(speechID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}

	c.Response.Status = 200
	return c.RenderJson(speech)
}

func (c SpeechController) Create() revel.Result {
	var (
		speech models.Speech
		err    error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&speech)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJson(errResp)
	}

	speech, err = models.AddSpeech(speech)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 201
	return c.RenderJson(speech)
}

func (c SpeechController) Update() revel.Result {
	var (
		speech models.Speech
		err    error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&speech)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	err = speech.UpdateSpeech()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	return c.RenderJson(speech)
}

func (c SpeechController) Delete(id string) revel.Result {
	var (
		err      error
		speech   models.Speech
		speechID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid speech id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	speechID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid speech id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	speech, err = models.GetSpeech(speechID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	err = speech.DeleteSpeech()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 204
	return c.RenderJson(nil)
}
