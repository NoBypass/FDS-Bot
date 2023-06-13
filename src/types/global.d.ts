export {}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      TOKEN: string
      API_URI: string
      API_VERSION: string
      CLIENT_ID: string
      ERROR_CHANNEL_ID: string
    }
  }
}