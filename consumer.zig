const std = @import("std");

extern fn somefunction() i64;

pub fn main() !void {}

export fn work() void {
    var x: i32 = 0;
    while (x < 100) {
        _ = somefunction();
        x += 1;
    }
}