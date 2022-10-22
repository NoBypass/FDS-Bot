import { Message, MessageEmbed } from "discord.js";
import fetch from "node-fetch";
import { hypixel_api_key, desc, client } from "./index";

export function createEmbed(content) {
    var channelname = content[1]
    var color = content[2]
    var title = content[3]
    var url = content[4]
    var authorname = content[5]
    var authoricon = content[6]
    var authorurl = content[7]
    var description = content[8]
    var thumbnail = content[9]
    var image = content[10]
    var field1name = content[11]
    var field1desc = content[12]
    var field1inline = content[13]
    var field2name = content[14]
    var field2desc = content[15]
    var field2inline = content[16]
    var field3name = content[17]
    var field3desc = content[18]
    var field3inline = content[19]
    var field4name = content[20]
    var field4desc = content[21]
    var field4inline = content[22]
    var field5name = content[23]
    var field5desc = content[24]
    var field5inline = content[25]
    var field6name = content[26]
    var field6desc = content[27]
    var field6inline = content[28]

    const newchannel = client.channels.cache.find(channel => (channel as any).name === channelname)
    const embed = new MessageEmbed()
        .setColor('#000000')
        .setTitle(title)
        .setURL(url)
        .setAuthor({ name: authorname, iconURL: authoricon, url: authorurl })
        .setDescription(description)
        .setThumbnail(thumbnail)
        .addFields(
            { name: field1name, value: field1desc, inline: field1inline },
            { name: field2name, value: field2desc, inline: field2inline },
            { name: field3name, value: field3desc, inline: field3inline },
            { name: field4name, value: field4desc, inline: field4inline },
            { name: field5name, value: field5desc, inline: field5inline },
            { name: field6name, value: field6desc, inline: field6inline },
        )
        .setImage(image)
        .setFooter({ text: desc });

        (newchannel as any).send({ embeds: [embed] });
}