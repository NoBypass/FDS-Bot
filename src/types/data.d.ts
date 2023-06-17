export type User = {
  id: string
  name: string
  password: string
  youtube?: string
  discord?: Owns
}

export type Owns = {
  id: string
  discordUser: DiscordUser
}

export type DiscordUser = {
  id: string
  uuid: string
  dailyStreak: number
  xp: number
  lastDailyAt: string
  minutesInVoice: number
  messagesSent: number
  registeredAt: string
  uploaded?: Uploaded[]
}

export type LinkedWith = {
  id: string
  at: string
  linkedByName: string
  linkedWith: DiscordUser
}

export type HypixelPlayer = {
  id: string
  isTracked: boolean
  linkedWith?: LinkedWith
}

export type Uploaded = {
  id: string
  at: string
  media: Media
}

export type Media = {
  id: string
  msg: string
  url: string
  likes: number
  dislikes: number
}

export type PlayedWith = {
  id: string
  hypixelPlayer: HypixelPlayer
}

export type MojangAccount = {
  id: string
  uuid: string
  name: string
  skin: string
  playedWith?: PlayedWith
}
