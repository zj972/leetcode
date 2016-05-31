//Distribution 46.77%,runtime 96ms
/**
 * @constructor
 */
var Queue = function() {
    this.queue = [];
};

/**
 * @param {number} x
 * @returns {void}
 */
Queue.prototype.push = function(x) {
    this.queue.push(x);
};

/**
 * @returns {void}
 */
Queue.prototype.pop = function() {
    this.queue.shift();
};

/**
 * @returns {number}
 */
Queue.prototype.peek = function() {
    return this.queue[0];
};

/**
 * @returns {boolean}
 */
Queue.prototype.empty = function() {
    return this.queue.length === 0;
};