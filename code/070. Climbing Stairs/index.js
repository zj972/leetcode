/**
 * @param {number} n
 * @return {number}
 */
//Distribution 52.82%,runtime 96ms
var climbStairs = function(n) {
    if (n <= 2) return n;
    var step = 3, s1 = 1, s2 = 2, tmp;
    while (step <= n) {
        tmp = s1 + s2;
        s1 = s2;
        s2 = tmp;
        ++step;
    }
    return s2;
};