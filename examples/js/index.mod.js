// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { fmt, parse, compile, evaluate } = require('./../../dist/dfl.mod.min.js')

console.log("************************************");
console.log();

const input = "@a + @b - @c";
const ctx = {a: 1, b: 4, c: 2};

console.log('Expression:');
console.log(input);
console.log();
console.log('Context:');
console.log(ctx);
console.log();

var { node, err } = parse(input);
if (err != null) {
  console.log('Error:');
  console.log(err);
  console.log();
  console.log("************************************");
} else {
  var { result, err } = evaluate(compile(node), ctx);
  console.log('Output:');
  console.log(result);
  console.log();
  if (err != null) {
    console.log('Error:');
    console.log(err);
    console.log();
  }
  console.log("************************************");
  console.log();
}
