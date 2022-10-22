import mongoose from 'mongoose'

const reqArray = {
    type: Array,
    required: true
}

const reqNum = {
    type: Number,
    required: true
}

const reqStr = {
    type: String,
    required: true
}

const tracked = new mongoose.Schema({
    uuid: reqStr,
    data: reqArray,
    date: reqArray,
    firsttracked: reqNum,
})

export default mongoose.model('tracked', tracked, 'tracked')