import mongoose from 'mongoose'

const reqString = {
    type: String,
    required: true
}

const reqNum = {
    type: Number,
    required: true
}

const field = {
    content: { type: Object },
    mode: { type: String }
}

const queue = new mongoose.Schema({
    activefields: { type: Number },
    fieldarr: [ field ]
})

export default mongoose.model('quueue', queue, 'queue')