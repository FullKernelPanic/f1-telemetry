const TelemetryType = {
    PUBLIC: "public",
    PRIVATE: "private"
}

/**
 * @typedef {id:Number, isAI:Boolean, isTeamMate: Boolean, name:String, nationality:String, networkId:Number, raceNumber: Number, teamId:Number, telemetryStatus:TelemetryType} Driver
 */


class State {
    /**
     * @private
     */
    _drivers;

    /**
     * @private
     */
    _driverPositions;

    constructor() {
        this._drivers = new Drivers();
        this._driverPositions = new DriverPositions(this._drivers);
    }

    /**
     * @returns {Drivers}
     */
    drivers() {
        return this._drivers;
    }

    /**
     * @returns {DriverPositions}
     */
    driverPositions() {
        return this._driverPositions;
    }
}

class DriverPositions {
    _data;

    constructor() {
        this._data = [];
    }

    update(data) {
        console.log('update');
        for (let i = 0; i < data.lapDatas.length; i++) {
            if (!this._data[i]) {
                this._data[i] = [data.lapDatas[i].carPosition];
            }

            this._data[i][data.lapDatas[i].currentLapNum] = data.lapDatas[i].carPosition;
        }
    }

    all() {
        return this._data;
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
