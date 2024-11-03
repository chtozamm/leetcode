export function rotateString(s: string, goal: string): boolean {
  if (s.length !== goal.length) return false
  const double = s + s
  return double.includes(goal)
}
