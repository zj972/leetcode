/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
//Distribution 23.58%,runtime 159ms
var isAnagram = function(s, t) {
    var sList = s.split("").sort().join();
    var tList = t.split("").sort().join();
    return (sList === tList);
};