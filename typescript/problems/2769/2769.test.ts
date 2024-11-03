import { theMaximumAchievableX } from "./2769"

describe("theMaximumAchievableX(num, t)", () => {
  it("should return the num incremented by 2 t times", () => {
    expect(theMaximumAchievableX(4, 1)).toBe(6)
    expect(theMaximumAchievableX(3, 2)).toBe(7)
  })
})
