/**
 * @param {string} val
 * @return {Object}
 */

var expect = function (val) {

    function throwErrorMsg(errorMsg) {
        throw new Error(errorMsg)
    }

    return {
        toBe: function toBe(checkValue) { return val === checkValue ? true : throwErrorMsg('Not Equal') },
        notToBe: function notToBe(checkValue) { return val !== checkValue ? true : throwErrorMsg('Equal') }
    }
}

/**
 * expect(5).toBe(5); // true
 * expect(5).notToBe(5); // throws "Equal"
 */