import mongoose from 'mongoose'

const reqString = {
  type: String,
  required: true,
}

const dailyRewards = new mongoose.Schema(
  {
    userId: reqString,
  },
  {
    timestamps: true,
  }
)

export default mongoose.model('daily-rewards', dailyRewards, 'daily-rewards')