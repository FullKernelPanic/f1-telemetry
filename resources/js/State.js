const TelemetryType = {
    PUBLIC: "public",
    PRIVATE: "private"
}

/**
 * @typedef {id:Number, isAI:Boolean, isTeamMate: Boolean, name:String, nationality:String, networkId:Number, raceNumber: Number, teamId:Number, telemetryStatus:TelemetryType} Driver
 */


class State {

    /**
     * @property {Driver[]} _drivers
     * @private
     */
    _drivers

    constructor() {
        this._drivers = new Drivers();
    }

    /**
     * @returns {Drivers}
     */
    drivers() {
        return this._drivers;
    }
}

class Drivers {
    /**
     * @private
     * @property {Driver[]} _drivers
     */
    _drivers

    /**
     * @private
     * @property Boolean _isValid
     */
    _isValid;

    constructor() {
        this._drivers = [];
        this._isValid = false;
    }

    /**
     * @param {Driver[]} drivers
     */
    update(drivers) {
        if (this._isChanged(drivers)) {
            this._drivers = drivers;
            this._isValid = false;
        }
    }

    /**
     * @returns {boolean}
     */
    isValid() {
        return this._isValid
    }

    /**
     * @returns {Driver[]}
     */
    all() {
        this._isValid = true;

        return this._drivers;
    }

    /**
     * @private
     * @param {Driver[]} drivers
     * @returns {boolean}
     */
    _isChanged(drivers) {
        if (this._drivers.length !== drivers.length) {
            return true;
        }

        for (let i = 0; i < drivers.length; i++) {
            if (this._drivers[i] !== drivers[i]) {
                return true;
            }
        }

        return false;
    }
}
