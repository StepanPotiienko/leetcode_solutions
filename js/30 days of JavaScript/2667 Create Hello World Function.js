/**
 * @return {Function}
 */
var createHelloWorld = function () {

    // Function should always return "Hello World",
    // no matter the input.
    return function (...args) {
        return "Hello World"
    }
};

/**
 * const f = createHelloWorld();
 * f(); // "Hello World"
 */