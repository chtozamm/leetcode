export function compressedString(word: string): string {
  const result: string[] = []
  let count = 1

  for (let i = 1; i <= word.length; i++) {
    let prev = word[i - 1]
    let curr = word[i]

    if (i < word.length && curr === prev) {
      count++
    } else {
      while (count > 0) {
        const chunk = Math.min(count, 9)
        result.push(chunk.toString(), prev)
        count -= chunk
      }
      count = 1
    }
  }

  return result.join("")
}
