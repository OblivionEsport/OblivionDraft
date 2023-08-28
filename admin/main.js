const teamNameSelectOne = document.getElementById('team1');
const teamNameSelectTwo = document.getElementById('team2');
const btnSelect = document.getElementById('btnSelect');

function changeSelect(select) {
    if (select.value === 'New') {
        emptyTeamSettings(select);
        btnSelect.disabled = true;
        return;
    }
    if (teamNameSelectOne.selectedIndex !== 0 && teamNameSelectTwo.selectedIndex !== 0 && teamNameSelectOne.selectedIndex !== teamNameSelectTwo.selectedIndex) {
        btnSelect.disabled = false;
    }
    loadTeamSettings(select, select.selectedIndex - 1);
}

fetch('/admin/teams/full')
    .then(response => response.json())
    .then(data => {
        data.forEach(team => {
            const option = document.createElement('option');
            option.value = team.name;
            option.text = team.name;
            teamNameSelectOne.appendChild(option);
            teamNameSelectTwo.appendChild(option.cloneNode(true));
        });
    });

async function saveTeamSettings(el) {
    const teamDiv = el.parentElement.parentElement.parentElement;
    const inputs = teamDiv.querySelectorAll('input');

    const newTeamData = {
        name: inputs[0].value,
        score: inputs[1].value,
        tag: inputs[2].value.toUpperCase(),
        color: inputs[3].value
    };
    let response;
    try {
        response = await fetch('/admin/teams/add/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newTeamData)
        });
    } catch (error) {
        console.error('An error occurred:', error);
    }
    if (response.ok) {
        console.log('Team added successfully!');
        // add an option to the select
        console.log(response.status);
        if (response.status == 201) {
            // new team
            const option = document.createElement('option');
            option.value = newTeamData.name;
            option.text = newTeamData.name;
            teamNameSelectOne.appendChild(option);
            teamNameSelectTwo.appendChild(option.cloneNode(true));
            el.value = newTeamData.name;
        }
    } else {
        console.error('Failed to add team.');
    }
}

function RemoveTeamSettings(el) {
    let selectObj = el.parentElement.parentElement.children[0];
    if (selectObj.value === 'New') {
        return;
    }
    fetch('/admin/teams/delete/' + selectObj.value)
        .then(resp => {
            if (resp.ok) {
                console.log('Team deleted successfully!');
                // remove the option from the select
                if (selectObj == teamNameSelectOne) {
                    teamNameSelectTwo.removeChild(teamNameSelectTwo.children[selectObj.selectedIndex]);
                    teamNameSelectOne.removeChild(teamNameSelectOne.children[selectObj.selectedIndex]);
                } else {
                    teamNameSelectOne.removeChild(teamNameSelectOne.children[selectObj.selectedIndex]);
                    teamNameSelectTwo.removeChild(teamNameSelectTwo.children[selectObj.selectedIndex]);
                }
                emptyTeamSettings(selectObj);
            } else {
                console.error('Failed to delete team.');
            }
        });
}

async function select() {
    // POST to /admin/teams/selected with the selected team
    resp = await fetch('/admin/teams/selected', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify([teamNameSelectOne.value, teamNameSelectTwo.value])
    });
    if (resp.ok) {
        console.log('Team selected successfully!');
    } else {
        console.error('Failed to select team.');
    }
}

function emptyTeamSettings(el) {
    parent = el.parentElement.parentElement;
    if (parent.className === "team-controls") {
        parent = parent.parentElement;
    }
    const inputs = parent.querySelectorAll('input');
    inputs.forEach(input => {
        input.value = '';
    });
}

function loadTeamSettings(el, i) {
    fetch('/admin/teams/full')
        .then(response => response.json())
        .then(data => {
            savedSettings = data[i]
            const inputs = el.parentElement.parentElement.querySelectorAll('input');
            inputs.forEach(input => {
                const id = input.id;
                if (savedSettings[id]) {
                    input.value = savedSettings[id];
                }
            });
        });

}