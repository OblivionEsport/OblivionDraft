const teamOneImg = document.querySelectorAll("img")[0];
const teamTwoImg = document.querySelectorAll("img")[1];


async function main() {
    let selected = await getTeamInfo();
    console.log(selected);
    teamOneImg.src = selected[0]["logoUrl"];
    teamTwoImg.src = selected[1]["logoUrl"];

}

async function getTeamInfo() {
    let resp = await fetch("/api/admin/teams/full")
    if (resp.status != 200) {
        return;
    }
    let data = await resp.json();
    let teams = data;
    resp = await fetch("/api/admin/teams/selected")
    if (resp.status != 200) {
        return;
    }
    data = await resp.json();
    let selectedTeam = data;
    let selected = []
    for (let i = 0; i < selectedTeam.length; i++) {
        for (let j = 0; j < teams.length; j++) {
            if (teams[j]["name"] == selectedTeam[i]) {
                selected[i] = teams[j];
            }
        }
    }
    return selected;
}

await main();