import DiscordJS, { ColorResolvable, Intents, Message, MessageEmbed, MessageActionRow, MessageButton, MessageAttachment, } from 'discord.js'
import mongoose from 'mongoose'
import 'dotenv/config'
import verifiedUsers from './schemas/verified-users'

const token = process.env.token
export const hypixel_api_key = process.env.hypixel_api_key
export const desc = 'Bot by NoBypass; v2.2.0; Last updated: 09.17.22; patch 5'
export const guildID = '62e15cc48ea8c9296133317f'

export const client = new DiscordJS.Client({
    intents: 32767,
})

export function delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}

export const modList = [
    '520968882077433856',
    '672835870080106509',
    '439955595878203403',
]
export const guild = client.guilds.cache.get('897642715368534027')

client.on('ready', async () => {
    await mongoose.connect(
        process.env.mongo_uri || '',
        {
            keepAlive: true
        })

    var duelsWins = await (await verifiedUsers.distinct('stats.duelswins').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    var duelsDeaths = await (await verifiedUsers.distinct('stats.duelsdeaths').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    var duelsKills = await (await verifiedUsers.distinct('stats.duelskills').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    var bridgeWins = await (await verifiedUsers.distinct('stats.bridgewins').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    var bedwarsFinals = await (await verifiedUsers.distinct('stats.bedwarsfinals').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    var bedwarsStars = await (await verifiedUsers.distinct('stats.bedwarsstars').exec()).map((x) => +x).reduce((partialSum, a) => partialSum + a, 0).toLocaleString()
    const stati = [
        'Do -help',
        'Duels Wins: ' + duelsWins,
        'Duels Deaths: ' + duelsDeaths,
        'Do -help',
        'Duels Kills: ' + duelsKills,
        'Bridge Wins: ' + bridgeWins,
        'Do -help',
        'BedWars Finals: ' + bedwarsFinals,
        'BedWars Stars: ' + bedwarsStars
    ]
    let i = 0
    const updateStatus = () => {
        client.user?.setPresence({
            status: 'online',
            activities: [
                {
                    name: stati[i]
                }
            ]
        })
        if (++i >= stati.length) {
            i = 0
        }
        setTimeout(updateStatus, 1000 * 5)
    }
    updateStatus()

    console.log('Bot started successfully.')
    let memberCount = require('./counters/member-count')
    let allApi = require('./counters/api-count')
    let handler = require('./command-handler')
    let textxp = require('./leveling/text-xp')
    let voicexp = require('./leveling/voice-xp')
    if (handler.default) handler = handler.default

    handler(client)
    memberCount(client)
    allApi(client)
    textxp(client)
    voicexp(client)
})

client.on('guildMemberAdd', member => {
    var memberRole = member.guild.roles.cache.find(role => role.name === "Guest");
    member.roles.add(memberRole);
});

client.on('interactionCreate', async (interaction) => {
    if (interaction.isButton()) {
        const gm = (await verifiedUsers.findOne({ memberid: interaction.user.id }).select('customstats.gm -_id') as any).customstats.gm
        if (interaction.customId === 'ann') {
            var newsletterName = 'announcement newsletter.'
            var rolel = '998241483159253044'
        } else if (interaction.customId === 'bot') {
            var newsletterName = 'bot update'
            var rolel = '998241874584293416'
        } else if (interaction.customId === 'tou') {
            var newsletterName = 'tournament'
            var rolel = '998242006109270077'
        } else if (interaction.customId === 'scr') {
            var newsletterName = 'scrims'
            var rolel = '998242097092104242'
        } else if (interaction.customId === 'bri') {
            var newsletterName = 'bridge queue'
            var rolel = '998242220920553473'
        } else if (interaction.customId === 'bed') {
            var newsletterName = 'bedwars queue'
            var rolel = '998242295268790353'
        } else if (interaction.customId === 'woo') {
            var newsletterName = 'woolgames queue'
            var rolel = '998242373257678968'
        } else if (interaction.customId === 'swe') {
            if (gm == true) return interaction.reply({ content: 'You cannot apply for the guild sorry.', ephemeral: true })
            var type = '〔📩〕sweat'
            createTicket(interaction, type)
        } else if (interaction.customId === 'gri') {
            if (gm == true) return interaction.reply({ content: 'You cannot apply for the guild sorry.', ephemeral: true })
            var type = '〔📩〕grinder'
            createTicket(interaction, type)
        } else if (interaction.customId === 'abs') {
            interaction.guild.channels.create(`〔📩〕misc-${interaction.user.tag}`, {
                parent: '1000781349248057354',
                topic: interaction.user.id,
                permissionOverwrites: [{
                    id: interaction.user.id,
                    allow: ['SEND_MESSAGES', 'VIEW_CHANNEL'],
                },
                {
                    id: interaction.guild.roles.everyone,
                    deny: ['VIEW_CHANNEL'],
                },
                ],
                type: "GUILD_TEXT",
            }).then(async c => {
                await verifiedUsers.findOneAndUpdate({ memberid: interaction.user.id }, {
                    form: {
                        cquestion: 1
                    }
                })
                interaction.reply({
                    content: `Created Ticket! <#${c.id}>`,
                    ephemeral: true
                })
                let m = await c.send('<@&897648604720795719>')
                m.delete()
                const buttonRow = new MessageActionRow().addComponents(
                    new MessageButton()
                        .setLabel('Delete Ticket')
                        .setStyle('DANGER')
                        .setCustomId('del'),
                )
                const startEmbed = new MessageEmbed()
                    .setColor('#000000')
                    .setDescription('You have created a "Misc" ticket. This should be used to share "private" information or questions to the server owners.')
                    .setFooter({ text: desc })

                await c.send({ embeds: [startEmbed], components: [buttonRow]  })
            })
        } else if (interaction.customId === 'den') {
            if (!modList.includes(interaction.user.id)) return interaction.reply({ content: 'You want to deny your own application?!? D:', ephemeral: true })
            interaction.channel.messages.fetch({ limit: 100 }).then(messages => {
                const secondMessage = messages.toJSON()[messages.size - 2];
                const memb = secondMessage.author.id;
                interaction.channel.send('<@' + memb + '> After reviewing your application, we sadly have to tell you that we have to deny it. If you **REALLY** want to join, we recomment becoming active in the community and playing with guild members. If they decide to vouch for you, you\'re almost guaranteed to be able to join!')
            }).catch(err => {
                console.log(err);
            });
        } else if (interaction.customId === 'apr') {
            if (!modList.includes(interaction.user.id)) return interaction.reply({ content: 'You can\'t approve yourself xD', ephemeral: true })
            interaction.channel.messages.fetch({ limit: 100 }).then(messages => {
                const secondMessage = messages.toJSON()[messages.size - 2];
                const memb = interaction.guild.members.cache.get(secondMessage.author.id)
                let role = memb.guild.roles.cache.find(r => r.id === "1001868724955009055")
                memb.roles.add(role)
                interaction.channel.send('<@' + memb + '> After reviewing your application, we are happy to welcome you to our guild! You have been given permission to view all the guild related channels. Whenever you\'re ready to join, message either a veteran, a mod or an admin to invite you and be on hypixel!')

            }).catch(err => {
                console.log(err);
            });
        } else if (interaction.customId === 'del') {
            interaction.reply({ content: 'Deleting channel in 5 seconds!', ephemeral: true })
            delay(5000).then(() => {
                interaction.channel.delete()
            })
        }

        if (interaction.customId === 'ann' || interaction.customId === 'bot' || interaction.customId === 'tou' || interaction.customId === 'scr' || interaction.customId === 'bri' || interaction.customId === 'bed' || interaction.customId === 'woo') {
            const member = await interaction.guild.members.fetch(interaction.user.id)
            if (member.roles.cache.has(rolel)) {
                await member.roles.remove(rolel)
                interaction.reply({ content: 'You have signed out of the **' + newsletterName + '** newsletter.', ephemeral: true })
            } else {
                await member.roles.add(rolel)
                interaction.reply({ content: 'You signed up for the **' + newsletterName + '** newsletter.', ephemeral: true })
            }
        }
    }
})

var questions = [
    'Hello! I have heard that you are interested in applying for the guild "FDS Employees". I have a few questions for you to answer. Please send your answer in one message and answer it honestly. \r First question: What is the gamemode do you play the most?',
    'What hobbies do you have? (IRL)',
    'Do you have any close friends within the guild?',
    'How long are you online for each day?',
    'Have you been banned from hypixel before and why? (Don\'t worry about telling me, I can keep a secret!)',
    'How much guild exp do you think you can get each day or week?',
    'When is your birthday? If you don\'t want to tell me that is absolutely fine.',
    'Okay, thank you for taking the time to answer all the questions. You will recieve a ping within 2 days telling you if you\'re in.',
]

async function createTicket(interaction, type) {
    interaction.guild.channels.create(`${type}-${interaction.user.tag}`, {
        parent: '1000781349248057354',
        topic: interaction.user.id,
        permissionOverwrites: [{
            id: interaction.user.id,
            allow: ['SEND_MESSAGES', 'VIEW_CHANNEL'],
        },
        {
            id: interaction.guild.roles.everyone,
            deny: ['VIEW_CHANNEL'],
        },
        ],
        type: "GUILD_TEXT",
    }).then(async c => {
        await verifiedUsers.findOneAndUpdate({ memberid: interaction.user.id }, {
            form: {
                cquestion: 1
            }
        })
        interaction.reply({
            content: `Created Ticket! <#${c.id}>`,
            ephemeral: true
        })
        const buttonRow = new MessageActionRow().addComponents(
            new MessageButton()
                .setLabel('Delete Ticket')
                .setStyle('DANGER')
                .setCustomId('del'),
        )
        const startEmbed = new MessageEmbed()
            .setColor('#000000')
            .setDescription(questions[0])
            .setFooter({ text: desc })

        await c.send({ embeds: [startEmbed], components: [buttonRow] })
    })
}

client.on('messageCreate', async (message) => {
    if (message.member.roles.cache.some(role => role.name === 'Guest')) return

    const messageChannel = message.channel as any
    if (messageChannel.parent.id === '1000781349248057354') {
      if (message.channel['name'].startsWith('〔📩〕misc')) return
        if (message.author.bot == true) return
        delay(1000).then(async () => {
            const currentQ = (await verifiedUsers.findOne({ memberid: message.member.id }).select('form.cquestion -_id') as any).form.cquestion
            await verifiedUsers.findOneAndUpdate({ memberid: message.member.id }, {
                form: {
                    cquestion: currentQ + 1,
                }
            })
            if (currentQ == questions.length - 1) {
                const emb = new MessageEmbed()
                    .setColor('#000000')
                    .setDescription(questions[questions.length - 1])
                    .setFooter({ text: desc })
                const buttonRow = new MessageActionRow().addComponents(
                    new MessageButton()
                        .setLabel('Approve')
                        .setStyle('SUCCESS')
                        .setCustomId('apr'),
                    new MessageButton()
                        .setLabel('Deny')
                        .setStyle('DANGER')
                        .setCustomId('den'),
                )
                message.reply({ content: '<@&897648604720795719>', embeds: [emb], components: [buttonRow] })
            } else if (currentQ - 1 < questions.length) {
                var emb = new MessageEmbed()
                    .setColor('#000000')
                    .setDescription(questions[currentQ])
                    .setFooter({ text: desc })
                await message.channel.send({ embeds: [emb] })
            }
        }).catch(err => { console.log(err) })
    }

    if (message.author.bot) return
    if (message.channel.id == '998259961102598205') {
        if (message.content.startsWith('<@&')) {
            if (message.content.length > 499) return message.reply('Your message was too long. Please write your message with less than 500 characters.')
            const ping = message.content.slice(0, 22)
            const args = message.content.slice(22, 500).split(', ')
            if (args.length > 2) return (message as any).reply('You used too many commas.')
            if (args.length < 1) return message.reply('You used no commas or have to specify more arguments.')

            const modeNum = message.content.slice(3, 21)
            if (modeNum == '998242097092104242') var mode = 'Scrims'
            else if (modeNum == '998242220920553473') mode = 'The Bridge'
            else if (modeNum == '998242295268790353') mode = 'BedWars'
            else if (modeNum == '998242373257678968') mode = 'WoolWars'
            else mode = 'with their Friend'

            const request = new MessageEmbed()
                .setColor('#000000')
                .setTitle(message.member.displayName + ' wants to queue ' + mode)
                .addFields(
                    { name: 'Submode', value: '``' + args[0] + '``', inline: false },
                    { name: 'Description', value: '``' + args[1] + '``', inline: false },
                )
                .setFooter({ text: desc });

            message.channel.send({ embeds: [request] });
            return message.delete()
        }
    }
});

client.login(token)
