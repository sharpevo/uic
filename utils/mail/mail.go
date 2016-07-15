package mail

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

func send_mail(subject, to, template_path string, values interface{}) error {

	RequestURI := "http://sendcloud.sohu.com/webapi/mail.send.json"
	body, err := compose_mail(to, template_path, values)
	if err != nil {
		return err
	}
	PostParams := url.Values{
		"api_user": {"iGeneTech_test_QmX8fK"},
		"api_key":  {"qEPgLrT05idBGq14"},
		"from":     {"support@igenetech.com"},
		"fromname": {"艾吉泰康"},
		"to":       {to},
		"subject":  {subject},
		"html":     {body},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		return err
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		return err
	}
	if is_sent, err := is_sent(BodyByte); !is_sent {
		return err
	}

	return err
}

func compose_mail(email, template_path string, values interface{}) (string, error) {

	tmpl, err := template.ParseFiles("utils/mail/mail_base.tpl", template_path)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	tmpl.ExecuteTemplate(writer, "base", values)
	writer.Flush()
	return b.String(), nil
}

func is_sent(res []byte) (bool, error) {
	response := make(map[string]interface{})
	_ = json.Unmarshal(res, &response)

	if response["message"] == "success" {
		return true, nil
	} else {
		return false,
			errors.New("SendCloud: " + fmt.Sprintf("%v", response["errors"]))
	}
}

const (
	RESET_SUBJECT  = "Password Reset Confirmation"
	RESET_TEMPLATE = "utils/mail/mail_reset_password.tpl"

	INVITE_SUBJECT  = "Invitation"
	INVITE_TEMPLATE = "utils/mail/mail_invitation.tpl"
)

type ResetArgs struct {
	ResetLink string
	Email     string
}

func SendResetMail(email, reset_link string) error {
	reset_args := ResetArgs{
		ResetLink: reset_link,
		Email:     email}
	err := send_mail(
		RESET_SUBJECT,
		email,
		RESET_TEMPLATE,
		reset_args)
	return err
}

type InvitationArgs struct {
	InviteCode string
	InviteLink string
}

func SendInvitation(email, invitation_code, invitation_link string) error {
	invitation := InvitationArgs{
		InviteCode: invitation_code,
		InviteLink: invitation_link}
	err := send_mail(
		INVITE_SUBJECT,
		email,
		INVITE_TEMPLATE,
		invitation)
	return err
}
