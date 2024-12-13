import { compressedString } from "./3163"

describe("compressedString", () => {
  it("correctly compress the string", () => {
    expect(compressedString("aaaaaaaaaaaaaabb")).toBe("9a5a2b")
    expect(compressedString("abcde")).toBe("1a1b1c1d1e")
    expect(compressedString("auuuuui")).toBe("1a5u1i")
    expect(compressedString("aaaaaaaaay")).toBe("9a1y")
  })
})
