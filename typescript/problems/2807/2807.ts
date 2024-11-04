import { ListNode } from "@/utils/linkedList"
import { greatestCommonDivisor } from "@/utils/math"

export function insertGreatestCommonDivisors(
  head: ListNode | null,
): ListNode | null {
  let curr = head
  while (curr && curr.next) {
    const temp = curr.next
    const gcd = greatestCommonDivisor(curr.val, curr.next.val)
    curr.next = new ListNode(gcd, temp)
    curr = curr.next.next
  }
  return head
}
