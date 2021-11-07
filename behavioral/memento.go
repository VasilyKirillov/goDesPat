package main

import "strings"

//memento
type TextWindowState struct {
	text string
}

func (s *TextWindowState) getText() string {
	return s.text
}

//originator
type TextWindow struct {
	currentText strings.Builder
}

func (t *TextWindow) addText(text string) {
	t.currentText.WriteString(text)
}
func (t *TextWindow) save() *TextWindowState {
	return &TextWindowState{text: t.currentText.String()}
}
func (t *TextWindow) restore(s *TextWindowState) {
	t.currentText = strings.Builder{}
	t.currentText.WriteString(s.getText())
}

//care-taker
type TextEditor struct {
	TextWindow
	savedTextWindowState *TextWindowState
}

func (e *TextEditor) hitSave() {
	e.savedTextWindowState = e.save()
}

func (e *TextEditor) hitUndo() {
	e.restore(e.savedTextWindowState)
}

func (e *TextEditor) write(text string) {
	e.addText(text)
}
func (e *TextEditor) print() string {
	return e.currentText.String()
}
