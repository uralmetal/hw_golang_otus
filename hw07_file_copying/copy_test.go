package main

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

const (
	lenTestString = 5000
)

func setUp(text []byte) (string, *os.File) {
	tempDir, err := os.MkdirTemp(".", "temp")
	if err != nil {
		panic(err)
	}
	tmpFile, err := os.CreateTemp(tempDir, "test")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	err = os.WriteFile(tmpFile.Name(), text, 0o644)
	if err != nil {
		panic(err)
	}
	return tempDir, tmpFile
}

func tearDown(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		panic(err)
	}
}

func generateText(length int) []byte {
	const russianAlphabet = "абвгдеёжзиклмопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИКЛМНОПРСТУФХЦЧШЩЭЮЯ"
	const englishAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digitals = "0123456789"
	const otherCharset = " 	\r\n\t"
	const charset = russianAlphabet + englishAlphabet + digitals + otherCharset

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return b
}

func TestCopy(t *testing.T) {
	randomText := generateText(lenTestString)
	testDir, testFile := setUp(randomText)
	defer tearDown(testDir)

	t.Run("simple case", func(t *testing.T) {
		testFileName := testFile.Name() + "_simple_case"
		err := Copy(testFile.Name(), testFileName, 0, 0)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText, copyContent)
	})

	t.Run("test limit 10", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_limit_10"
		testLimit := int64(10)
		err := Copy(testFile.Name(), testFileName, 0, testLimit)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText[:testLimit], copyContent)
	})

	t.Run("test limit 100", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_limit_10"
		testLimit := int64(100)
		err := Copy(testFile.Name(), testFileName, 0, testLimit)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText[:testLimit+1], copyContent)
	})

	t.Run("test overwrite file", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_overwrite"
		testLimit := int64(100)
		err := Copy(testFile.Name(), testFileName, 0, testLimit)
		require.NoError(t, err)
		err = Copy(testFile.Name(), testFileName, 0, 0)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText, copyContent)
	})

	t.Run("test offset", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_offset"
		testOffset := int64(100)
		err := Copy(testFile.Name(), testFileName, testOffset, 0)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText[testOffset:], copyContent)
	})

	t.Run("test offset and limit", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_limit_offset"
		testOffset := int64(100)
		testLimit := int64(100)
		err := Copy(testFile.Name(), testFileName, testOffset, testLimit)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText[testOffset:testOffset+testLimit], copyContent)
	})

	t.Run("test with offset more file", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_long_offset"
		testOffset := int64(lenTestString + 100)
		err := Copy(testFile.Name(), testFileName, testOffset, 0)
		require.Error(t, err)
		_, err = os.ReadFile(testFileName)
		require.Error(t, err)
	})

	t.Run("test with limit more file", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_long_limit"
		testLimit := int64(lenTestString + 100)
		err := Copy(testFile.Name(), testFileName, 0, testLimit)
		require.NoError(t, err)
		copyContent, err := os.ReadFile(testFileName)
		require.NoError(t, err)
		require.Equal(t, randomText, copyContent)
	})

	t.Run("test with limit more file", func(t *testing.T) {
		testFileName := testFile.Name() + "_test_long_limit"
		err := Copy(testDir, testFileName, 0, 0)
		require.Error(t, err)
	})

	t.Run("invalid parameters", func(t *testing.T) {
		testFileName := testFile.Name() + "_invalid_params"
		require.Error(t, Copy(testFile.Name()+"_invalid", testFileName, 0, 0))
		require.Error(t, Copy(testFile.Name(), testFileName, -1, 0))
		require.Error(t, Copy(testFile.Name(), testFileName, 0, -1))
	})
}
