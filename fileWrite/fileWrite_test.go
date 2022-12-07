package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("ueharayuto")
    want := "Hi, ueharayuto. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}