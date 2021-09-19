const app = new Vue({
    el: '#app',
    data: {
        notes: ["A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"],
        scales: ["minor", "major", "harmonic minor"],
        rootNote: "A",
        scale: "minor",
        tuning: "E A D G B E",
    },
    computed: {
        imageSRC: function () {
            const root = encodeURIComponent(this.rootNote)
            const scale = encodeURIComponent(this.scale)
            const tuning = encodeURIComponent(this.tuning)
            return "/scale?root=" + root + "&type=" + scale + "&tuning=" + tuning;
        },
    }
})