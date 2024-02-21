package mock

import (
	"errors"
	"testing"
)

// Dummy objects are passed around but never actually used. Usually they are just used to fill parameter lists.
// 需要傳遞參數，但知道參數不會被用到，只是為了符合測試的參數需求，不影響其他邏輯的執行
// 目的 : 測試MessageService有沒有成功建立
type DummyNotifier struct{}

func (d DummyNotifier) Send(message string) error {
	return nil
}

func TestMessageService_Dummy(t *testing.T) {
	service := NewMessageService(DummyNotifier{})
	if service == nil {
		t.Errorf("Something wrong")
	}
}

// Fake objects actually have working implementations, but usually take some shortcut which makes them not suitable for production (an InMemoryTestDatabase is a good example).
// 會有一些簡單的邏輯或行為來模擬真實對象的功能，但不會涉及完整的production邏輯
// 目的 : 測試消息是不是加薪
type FakeNotifier struct{}

func (f FakeNotifier) Send(message string) error {
	if message == "加薪" {
		return errors.New("invalid sentence")
	}
	return nil
}

func TestMessageService_Fake(t *testing.T) {
	messageService := NewMessageService(&FakeNotifier{})

	err := messageService.SendAlert("加薪")
	if err == nil {
		t.Error("這不是一個加薪的請求，耶")
	}
}

// Stubs provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.
// 返回預先定義好的canned answers，主要是提供測試中所需要假的數據。 StubNotifier用以取代原本相依的第三方元件(可能一些緩存,對列服務)，讓測試時不用擔心第三方的邏輯會導致測試失敗。
// 目的 : 驗證 MessageService 在發送成功或失敗的 Stub 條件下,MessageService還是可以如預期的執行SendAlert方法
type StubNotifier struct {
	success bool
}

func (s StubNotifier) Send(message string) error {
	if !s.success {
		return errors.New("警告失敗")
	}
	return nil
}

func TestMessageService_Stub_Success(t *testing.T) {
	messageService := NewMessageService(&StubNotifier{success: true})

	err := messageService.SendAlert("任何句子")

	if err != nil {
		t.Errorf("預計成功寄出警告消息, got %v", err)
	}
}

func TestMessageService_Stub_Fail(t *testing.T) {
	messageService := NewMessageService(&StubNotifier{success: false})

	err := messageService.SendAlert("任何句子")

	if err == nil {
		t.Error("預計警告失敗，但出現錯誤")
	}
}

// Spies are stubs that also record some information based on how they were called. One form of this might be an email service that records how many messages it was sent.
// called 字段用於記錄 SendAlert 方法是否被調用過。
// sentence 字段則記錄調用 SendAlert 方法時傳遞的句子。
// Spy 預期是「某些互動是否發生過」這類確認測試。
// 目的: 通過 Spy 驗證 MessageService 中與 Notifier 的互動是否正確,包括調用次數和參數傳遞。
type SpyNotifier struct {
	called   bool
	sentence string
}

func (s *SpyNotifier) Send(message string) error {
	s.called = true
	s.sentence = message
	return nil
}

func TestMessageService_Spy(t *testing.T) {
	spyNotifier := &SpyNotifier{}
	messageService := NewMessageService(spyNotifier)

	message := "喜歡上班"
	_ = messageService.SendAlert(message)

	if !spyNotifier.called {
		t.Error("預期spyNotifier被呼叫")
	}

	if spyNotifier.sentence == "不喜歡上班" {
		t.Errorf("老闆難過QQ")
	}

}

// Mocks are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.
// 要用於模擬複雜的互動和行為，並驗證這些互動是否按照預期發生
// Mock預期更多是事先定義「這些互動次數必須完全正確」,以實現規範測試的效果。
type MockNotifier struct {
	calls []string
}

func (m *MockNotifier) Send(message string) error {
	m.calls = append(m.calls, "Send方法")

	return nil
}

func (m *MockNotifier) VerifyCalls(expectedCalls []string) bool {
	if len(expectedCalls) != len(m.calls) {
		return false
	}
	for i, call := range expectedCalls {
		if call != m.calls[i] {
			return false
		}
	}
	return true
}

func  TestMessageService_Mock(t *testing.T) {
	mockNotifier := &MockNotifier{}
	messageService := NewMessageService(mockNotifier)

	message := "喜歡上班"
	_ = messageService.SendAlert(message)

	expectedCalls := []string{"Send方法"}
	if !mockNotifier.VerifyCalls(expectedCalls) {
		t.Errorf("預期呼叫 %v, 但得到 %v", expectedCalls, mockNotifier.calls)
	}
}
