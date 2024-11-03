import { insertGreatestCommonDivisors } from "./2807";
import { createLinkedList } from "@/utils/linkedList";

describe("insertGreatestCommonDivisors", () => {
  it("should insert the greatest common divisor between each pair of nodes", () => {
    const before = createLinkedList([18, 6, 10, 3]);
    const after = createLinkedList([18, 6, 6, 2, 10, 1, 3]);
    expect(insertGreatestCommonDivisors(before)).toEqual(after);
  });

  it("should return the same list when there is only one node", () => {
    const before = createLinkedList([7]);
    const after = createLinkedList([7]);
    expect(insertGreatestCommonDivisors(before)).toEqual(after);
  });

  it("should return null for an empty list", () => {
    const before = createLinkedList([]);
    const after = createLinkedList([]);
    expect(insertGreatestCommonDivisors(before)).toEqual(after);
  });

  it("should handle a list with two nodes", () => {
    const before = createLinkedList([8, 4]);
    const after = createLinkedList([8, 4, 4]);
    expect(insertGreatestCommonDivisors(before)).toEqual(after);
  });
});
