export const getNeededXp = (level: number): number => {
  if (level < 10) return Math.pow(level, 2) * 100
  return 10000
}

export const getTotalXp = (level: number, xp: number) => {
  for (let i = 1; i < level; i++) {
    xp += getNeededXp(i)
  }
  return xp
}
