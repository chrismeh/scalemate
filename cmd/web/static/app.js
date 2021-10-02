function sendScaleRequest(updateChordSelector) {
    const root = encodeURIComponent(document.getElementById("root").value);
    const scale = encodeURIComponent(document.getElementById("scale").value);
    const tuning = encodeURIComponent(document.getElementById("tuning").value);
    const frets = encodeURIComponent(document.getElementById("frets").value);

    let url = `/api/scale?root=${root}&type=${scale}&tuning=${tuning}&frets=${frets}`

    let chord = document.getElementById("chord").value
    if (chord !== "-") {
        url += "&chord=" + encodeURIComponent(chord);
    }

    fetch(url)
        .then(resp => resp.json())
        .then(json => {
            document.getElementById("scale-image").src = `data:image/png;base64,${json.picture}`;

            if (!updateChordSelector) {
                return;
            }

            let chordSelect = document.getElementById("chord")
            while (chordSelect.options.length > 0) {
                chordSelect.remove(0);
            }

            let emptyOption = document.createElement("option");
            emptyOption.innerHTML = "-";
            chordSelect.appendChild(emptyOption);

            for (let chord of json.chords) {
                let opt = document.createElement("option")
                opt.value = chord;
                opt.innerHTML = chord;
                chordSelect.appendChild(opt)
            }
        })
}