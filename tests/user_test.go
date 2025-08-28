package integration_test

import (
	"encoding/json"
	"hulio-user-service/handler/response"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestUserAPI(t *testing.T) {
	//1.login with not-existed username
	millis := time.Now().UnixMilli()
	msStr := strconv.FormatInt(millis, 10)
	username := "user" + msStr
	password := "pwd" + msStr
	res := login(t, username, password)
	if res.Code != 001 {
		t.Fatalf("expected user not exist, got message %s", res.Message)
	}

	//2.register user
	regRes := register(t, username, password)
	if regRes.Success != true {
		t.Fatalf("expected register success, got message %s", regRes.Message)
	}

	//3.register user with same username again
	regRes2 := register(t, username, password)
	if regRes2.Code != 003 {
		t.Fatalf("expected username existed, got message %s", regRes.Message)
	}

	//4.login with error password
	res2 := login(t, username, "password")
	if res2.Code != 002 {
		t.Fatalf("expected password error, got message %s", res.Message)
	}

	//5.login with correct password
	res3 := login(t, username, password)
	if res3.Success != true {
		t.Fatalf("expected login success, got message %s", res.Message)
	}

	// extract token from login response
	token, ok := extractToken(t, res3)
	if !ok || token == "" {
		t.Fatalf("expected token in login response")
	}

	//6.get user info without token
	noTokenRes := getUserInfo(t, "")
	if noTokenRes.Code != 001 {
		t.Fatalf("expected failure without token, got message %s", noTokenRes.Message)
	}

	//7.get user info with error token
	badTokenRes := getUserInfo(t, "invalid-token")
	if badTokenRes.Success != false {
		t.Fatalf("expected failure with invalid token, got success")
	}

	//8.get user info
	infoRes := getUserInfo(t, token)
	if infoRes.Success != true {
		t.Fatalf("expected get user info success, got message %s", infoRes.Message)
	}

	//9.edit user info: phone avator email and nickname
	newNickname := "nick_" + msStr
	newAvatar := "https://example.com/avatar/" + msStr + ".png"
	newEmail := "u" + msStr + "@example.com"
	newPhone := "138" + msStr[len(msStr)-8:]
	newExtinfo := "v" + msStr[len(msStr)-3:]
	updRes := updateUser(t, token, newNickname, newAvatar, newEmail, newPhone, newExtinfo)
	if updRes.Success != true {
		t.Fatalf("expected update user success, got message %s", updRes.Message)
	}

	//10.get user info after update
	infoRes2 := getUserInfo(t, token)
	if infoRes2.Success != true {
		t.Fatalf("expected get user info success after update, got message %s", infoRes2.Message)
	}
	// verify updated fields
	userBytes, _ := json.Marshal(infoRes2.Data)
	var parsed struct {
		User struct {
			Nickname string `json:"nickname"`
			Avatar   string `json:"avatar"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			Extinfo  string `json:"extinfo"`
		} `json:"user"`
	}
	_ = json.Unmarshal(userBytes, &parsed)
	if parsed.User.Nickname != newNickname || parsed.User.Avatar != newAvatar || parsed.User.Email != newEmail || parsed.User.Phone != newPhone || parsed.User.Extinfo != newExtinfo {
		t.Fatalf("user fields not updated as expected")
	}

	//11.delete user
	delRes := deleteUser(t, token)
	if delRes.Success != true {
		t.Fatalf("expected delete user success, got message %s", delRes.Message)
	}

	//12.login again, expect user not exist
	res4 := login(t, username, password)
	if res4.Code != 001 {
		t.Fatalf("expected user not exist after deletion, got message %s", res4.Message)
	}
}

func login(t *testing.T, username string, password string) response.Response {
	return doJSON(t, "login", http.MethodPost, "http://localhost:8080/api/v1/user/login", "", httpPayload{
		"username": username,
		"password": password,
	})
}

func register(t *testing.T, username string, password string) response.Response {
	return doJSON(t, "register", http.MethodPost, "http://localhost:8080/api/v1/user/register", "", httpPayload{
		"username": username,
		"password": password,
	})
}

// helper to extract token from login response.Data
func extractToken(t *testing.T, res response.Response) (string, bool) {
	// Data is decoded as map[string]interface{}
	bytesData, err := json.Marshal(res.Data)
	if err != nil {
		t.Fatalf("marshal data failed: %v", err)
	}
	var parsed struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(bytesData, &parsed); err != nil {
		t.Fatalf("unmarshal token failed: %v", err)
	}
	return parsed.Token, parsed.Token != ""
}

func getUserInfo(t *testing.T, token string) response.Response {
	return doJSON(t, "get user info", http.MethodGet, "http://localhost:8080/api/v1/user", token, nil)
}

func updateUser(t *testing.T, token, nickname, avatar, email, phone, extinfo string) response.Response {
	return doJSON(t, "update user", http.MethodPut, "http://localhost:8080/api/v1/user", token, httpPayload{
		"nickname": nickname,
		"avatar":   avatar,
		"email":    email,
		"phone":    phone,
		"extinfo":  extinfo,
	})
}

func deleteUser(t *testing.T, token string) response.Response {
	return doJSON(t, "delete user", http.MethodDelete, "http://localhost:8080/api/v1/user", token, nil)
}

type httpPayload map[string]string

func doJSON(t *testing.T, label, method, url, token string, payload httpPayload) response.Response {
	t.Helper()
	var bodyReader *strings.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("marshal payload failed: %v", err)
		}
		bodyReader = strings.NewReader(string(b))
	} else {
		bodyReader = strings.NewReader("")
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatalf("build request failed: %v", err)
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var res response.Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}
	t.Logf("%s res: %+v", label, res)
	return res
}
