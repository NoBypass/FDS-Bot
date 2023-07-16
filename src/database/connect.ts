import { MainModel } from './schema'
import { log } from 'console'
import chalk from 'chalk'
import { config } from 'dotenv'
import mongoose from 'mongoose'

const connectToDB = async () => {
  config()
  const uri = `mongodb+srv://${process.env.DB_USERNAME}:${process.env.DB_PASSWORD}@cluster0.ue5hw.mongodb.net/${process.env.DB_NAME}?retryWrites=true&w=majority`

  try {
    await mongoose.connect(uri)
    const startTime = Date.now()
    await MainModel.findOne({ ign: 'NoBypass' })
    log(
      chalk.green(
        `Successfully connected to database (completed in ${
          Date.now() - startTime
        }ms)`,
      ),
    )
  } catch (e) {
    log(chalk.red('Failed to connect to database, error: '))
    throw new Error(e)
  }
}

export default connectToDB
