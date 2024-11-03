import { rotateString } from "./796"

describe("rotateString", () => {
  describe("valid rotations", () => {
    it("should return true for valid rotations", () => {
      expect(rotateString("abcde", "cdeab")).toBe(true)
      expect(rotateString("hello", "llohe")).toBe(true)
      expect(rotateString("", "")).toBe(true)
    })
  })

  describe("invalid rotations", () => {
    it("should return false for invalid rotations", () => {
      expect(rotateString("abcde", "abced")).toBe(false)
      expect(rotateString("abc", "def")).toBe(false)
      expect(rotateString("abc", "bac")).toBe(false)
      expect(rotateString("abc", "ab")).toBe(false)
      expect(rotateString("", "a")).toBe(false)
    })
  })
})
