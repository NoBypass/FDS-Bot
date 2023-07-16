import { client } from '../index'
import { JwtResponse, MojangAccountWithHypixelPlayer } from '../types/api'
import { HypixelPlayer, MojangAccount, PlayedWith } from '../types/data'

const { API_URI, API_VERSION } = process.env

const graphqlCall = async <T>(
  query: string,
  schema: string,
): Promise<Partial<T>> => {
  return await fetch(`${API_URI}/${API_VERSION}/graphql`, {
    method: 'POST',
    body: `${query} {\n${schema}}`,
    headers: {
      'Content-Type': 'application/graphql',
      Authorization: `Bearer ${client.auth.token}`,
    },
  }).then((res) => res.json())
}

const postCall = async (body: string, route: string): Promise<any> => {
  return await fetch(`${API_URI}/${API_VERSION}/${route}`, {
    method: 'POST',
    body: body,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${client.auth.token}`,
    },
  }).then((res) => res.json())
}

export const api = {
  get: {
    mojangAccountByName: async (name: string, schema: string) => {
      return await graphqlCall(`mojangAccount(name: ${name})`, schema)
    },
  },
  add: {
    mojangAccount: async (name: string, schema: string) => {
      return await graphqlCall<MojangAccount>(
        `addMojangAccount(name: ${name})`,
        schema,
      )
    },
    hypixelPlayer: async (schema: string) => {
      return await graphqlCall<HypixelPlayer>(
        'addHypixelPlayer(isTracked: false)',
        schema,
      )
    },
  },
  connect: {
    mojangAccountWithHypixelPlayer: async (
      { mojangAccountId, hypixelPlayerId }: MojangAccountWithHypixelPlayer,
      schema: string,
    ) => {
      return await graphqlCall<PlayedWith>(
        `connectMojangAccountWithHypixelPlayer(
                    mojangAccountId: ${mojangAccountId},
                    hypixelPlayerId: ${hypixelPlayerId}
                )`,
        schema,
      )
    },
  },
}

export const login = async (
  name: string,
  password: string,
): Promise<JwtResponse> => {
  const loginAttempt = await postCall(
    JSON.stringify({ name, password }),
    'login',
  )
  return loginAttempt.jwt != null
    ? loginAttempt
    : await postCall(JSON.stringify({ name, password }), 'register')
}
