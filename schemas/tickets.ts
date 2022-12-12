import mongoose from 'mongoose'

const reqString = {
    type: String,
    required: true
}

const reqNum = {
    type: Number,
    required: true
}

const tickets = new mongoose.Schema({
    memberid: reqString,
    createdAt: reqNum,
    closedAt: reqNum,
    answers: [reqString],
    type: reqString,
    questionsAsked: reqNum,
})

export default mongoose.model('tickets', tickets, 'tickets')