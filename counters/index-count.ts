export async function getIndex(data) {
    var b1 = Math.sqrt(2); var a1 = Math.sqrt((data.networkExp || 0) + 15312.5); var c1 = 25 * b1; var d1 = 125 / b1; var e1 = a1 - d1; var networklevel = e1 / c1; var nwlvlrounded = Math.floor(networklevel)

    var swlevel = (data.player?.stats?.SkyWars?.skywars_experience || 0 - 15000) / 10000 + 12

    var bridgelosses = data.player?.stats?.Duels?.bridge_duel_losses || 0 + data.player?.stats?.Duels?.bridge_doubles_losses || 0 + data.player?.stats?.Duels?.bridge_threes_losses || 0 + data.player?.stats?.Duels?.bridge_four_losses || 0 + data.player?.stats?.Duels?.bridge_3v3v3v3_losses || 0 + data.player?.stats?.Duels?.bridge_2v2v2v2_losses || 0 + data.player?.stats?.Duels?.capture_threes_losses || 0

    var rduelwins = data.player?.stats?.Duels?.wins || 0 - data.player?.achievements?.duels_bridge_wins || 0

    var duelwlr = (data.player?.stats?.Duels?.wins || 0) / (data.player?.stats?.Duels?.losses || 0)
    var wweffkd = ((data.player?.stats?.WoolGames?.wool_wars?.stats?.kills || 0) + (data.player?.stats?.WoolGames?.wool_wars?.stats?.assists || 0 / 2)) / (data.player?.stats?.WoolGames?.wool_wars?.stats?.deaths || 0)
    var bridgewlr = (data.player?.achievements?.duels_bridge_wins || 0) / bridgelosses
    var swkd = (data.player?.stats?.SkyWars?.kills || 0) / (data.player?.stats?.SkyWars?.deaths || 0)
    var bwfkdr = (data.player?.stats?.Bedwars?.final_kills_bedwars || 0) / (data.player?.stats?.Bedwars?.final_deaths_bedwars || 0)

    var bwlevel = getBedWarsStar(data.player?.stats?.Bedwars?.Experience || 0)
    var wglevel = getWoolWarsStar(data.player?.tats?.WoolGames?.progression?.experience || 0) + 1

    var bwindex = parseInt(Math.round(bwlevel * 15 + Math.pow(bwfkdr, 2.5) * 3.5).toString().slice(0, -2)) || 0
    var duelindex = parseInt(Math.round(rduelwins / 3.5 + Math.pow(duelwlr, 1.5)).toString().slice(0, -2)) || 0
    var swindex = parseInt(Math.round(Math.pow(swlevel, 2.3) + Math.pow(swkd, 5)).toString().slice(0, -2)) || 0
    var bridgeindex = parseInt(Math.round(data.player?.achievements?.duels_bridge_wins || 0 + Math.pow(bridgewlr, 1.7)).toString().slice(0, -2)) || 0
    var miscindex = parseInt(Math.round((data.player?.achievementPoints || 0) * 2.2 + nwlvlrounded).toString().slice(0, -2)) || 0
    var wwindex = parseInt(Math.round(wglevel * 15 + Math.pow(wweffkd, 3.5) * 3.5).toString().slice(0, -2)) || 0

    var indexarr = [bwindex, bridgeindex, miscindex, wwindex, swindex, duelindex]
    for (let i = 0; i < indexarr.length; i++) {
        indexarr[i] = Math.pow(Math.sqrt(indexarr[i]), 1.85)
        indexarr[i] = Math.round(indexarr[i])
        var index = 0
        indexarr.forEach(item => {
            index += item;
        });
    }

    const indexes = {
        indexarr: indexarr, 
        index: index
    }

    return indexes
};

const EASY_LEVELS = 4;
const EASY_LEVELS_XP = 7000;
const XP_PER_PRESTIGE = 96 * 5000 + EASY_LEVELS_XP;
const LEVELS_PER_PRESTIGE = 100;
const HIGHEST_PRESTIGE = 10;
function getExpForLevel(level) {
    if (level == 0) return 0;
    var respectedLevel = getLevelRespectingPrestige(level);
    if (respectedLevel > EASY_LEVELS) {
        return 5000;
    }
    switch (respectedLevel) {
        case 1:
            return 1000;
        case 2:
            return 2000;
        case 3:
            return 3000;
        case 4:
            return 4000;
    }
    return 5000;
}
function getLevelRespectingPrestige(level) {
    if (level > HIGHEST_PRESTIGE * LEVELS_PER_PRESTIGE) {
        return level - HIGHEST_PRESTIGE * LEVELS_PER_PRESTIGE;
    }
    else {
        return level % LEVELS_PER_PRESTIGE;
    }
}
export function getWoolWarsStar(exp) {
    var prestiges = Math.floor(exp / XP_PER_PRESTIGE);
    var level = prestiges * LEVELS_PER_PRESTIGE;
    var expWithoutPrestiges = exp - (prestiges * XP_PER_PRESTIGE);
    for (let i = 1; i <= EASY_LEVELS; ++i) {
        var expForEasyLevel = getExpForLevel(i);
        if (expWithoutPrestiges < expForEasyLevel) {
            break;
        }
        level++;
        expWithoutPrestiges -= expForEasyLevel;
    }
    return level + Math.floor(expWithoutPrestiges / 5000);
}

const EASY_LEVELSBW = 4;
const EASY_LEVELS_XPBW = 7000;
const XP_PER_PRESTIGEBW = 96 * 5000 + EASY_LEVELS_XPBW;
const LEVELS_PER_PRESTIGEBW = 100;
const HIGHEST_PRESTIGEBW = 10;
function getbwExpForLevel(level) {
    if (level == 0) return 0;
    var respectedLevel = getbwLevelRespectingPrestige(level);
    if (respectedLevel > EASY_LEVELSBW) {
        return 5000;
    }
    switch (respectedLevel) {
        case 1:
            return 500;
        case 2:
            return 1000;
        case 3:
            return 2000;
        case 4:
            return 3500;
    }
    return 5000;
}
function getbwLevelRespectingPrestige(level) {
    if (level > HIGHEST_PRESTIGEBW * LEVELS_PER_PRESTIGEBW) {
        return level - HIGHEST_PRESTIGEBW * LEVELS_PER_PRESTIGEBW;
    }
    else {
        return level % LEVELS_PER_PRESTIGEBW;
    }
}
export function getBedWarsStar(exp) {
    var prestiges = Math.floor(exp / XP_PER_PRESTIGEBW);
    var level = prestiges * LEVELS_PER_PRESTIGEBW;
    var expWithoutPrestiges = exp - (prestiges * XP_PER_PRESTIGEBW);
    for (let i = 1; i <= EASY_LEVELSBW; ++i) {
        var expForEasyLevel = getbwExpForLevel(i);
        if (expWithoutPrestiges < expForEasyLevel) {
            break;
        }
        level++;
        expWithoutPrestiges -= expForEasyLevel;
    }
    return level + Math.floor(expWithoutPrestiges / 5000);
}