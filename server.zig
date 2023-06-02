const std = @import("std");

pub fn main() !void {}

var _counter: i64 = 0;

export fn somefunction() i64 {
    _counter += 1;
    return _counter;
}

export fn counter() i64 {
    return _counter;
}