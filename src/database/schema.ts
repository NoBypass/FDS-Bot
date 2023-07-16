import mongoose, { Document, Schema, Types } from 'mongoose'

const reqNum = {
  type: Number,
  required: true,
}

const reqString = {
  type: String,
  required: true,
}

const verifiedUsersSchema = new Schema({
  ign: reqString,
  uuid: reqString,
  memberid: reqString,
  form: {
    cquestion: {
      type: Number,
      required: false,
      default: 0,
    },
  },
  xp: {
    type: Number,
    default: 0,
  },
  level: {
    type: Number,
    default: 1,
  },
  customtag: String,
  lastclaimed: reqNum,
  streak: {
    type: Number,
    default: 0,
  },
})

interface IVerifiedUser extends Document {
  ign: string
  uuid: string
  memberid: string
  xp: number
  level: number
  customtag?: string
  lastclaimed: number
  streak: number
}

export const MainModel = mongoose.model<IVerifiedUser>(
  'verifiedUsers',
  verifiedUsersSchema,
  'verifiedUsers',
)
