class StartLights extends HTMLElement {
    static observedAttributes = ["lights"];

    constructor() {
        super();

        const shadow = this.attachShadow({ mode: "open" });
        let container = document.createElement("div");
        container.innerHTML = `
            <style>
            .off { fill: #363636; }
            .on {fill: #EE1111; }

            </style>
            <div class="container">
                <svg width="40px" height="40px"><circle cx="20px" cy="20px" r="16px" class="off"></circle></svg>
                <svg width="40px" height="40px"><circle cx="20px" cy="20px" r="16px" class="off"></circle></svg>
                <svg width="40px" height="40px"><circle cx="20px" cy="20px" r="16px" class="off"></circle></svg>
                <svg width="40px" height="40px"><circle cx="20px" cy="20px" r="16px" class="off"></circle></svg>
                <svg width="40px" height="40px"><circle cx="20px" cy="20px" r="16px" class="off"></circle></svg>
            </div>
        `;

        shadow.appendChild(container);
    }

    connectedCallback() {
        console.log("Custom element added to page.");
    }

    disconnectedCallback() {
        console.log("Custom element removed from page.");
    }

    adoptedCallback() {
        console.log("Custom element moved to new page.");
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (name === "lights") {
            let lamps = this.shadowRoot.querySelectorAll("circle");

            for (let i = 0; i < 5; i++) {
                if (i <= newValue - 1) {
                    lamps[i].classList.remove("off");
                    lamps[i].classList.add("on");
                } else {
                    lamps[i].classList.remove("on");
                    lamps[i].classList.add("off");
                }
            }
        }
    }
}

customElements.define("start-lights", StartLights);
