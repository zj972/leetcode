/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    var sList = s.split("").sort().join();
    var tList = t.split("").sort().join();
    return (sList === tList);
};