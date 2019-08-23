// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { parse, compile, evaluate, fmt } = global.dfl;

function log(str) {
  console.log(str.replace(/\n/g, "\\n").replace(/\t/g, "\\t").replace(/"/g, "\\\""));
}

describe('dfl', () => {

  it('fmt', () => {
    var { expression, remainder, err } = fmt("@a + @b - @c")
    expect(err).toBeNull();
    expect(remainder).toEqual("");
    expect(expression).toEqual("(@a + (@b - @c))");
  });

  it('parse', () => {
    var { node, remainder, err } = parse("@a + @b - @c")
    expect(err).toBeNull();
    expect(remainder).toEqual("");
    expect(node).toBeDefined();
    expect(Object.keys(node).sort()).toEqual(["Compile", "Evaluate", "Pretty", "__internal_object__",]);
  });

  it('compile', () => {
    var { node, remainder, err } = parse("@a + @b - @c")
    expect(err).toBeNull();
    expect(remainder).toEqual("");
    expect(node).toBeDefined();
    expect(Object.keys(node).sort()).toEqual(["Compile", "Evaluate", "Pretty", "__internal_object__",]);
    // test method
    expect(node.Compile()).toBeDefined();
    expect(Object.keys(node.Compile()).sort()).toEqual(["Compile", "Evaluate", "Pretty", "__internal_object__",]);
    // test function
    expect(compile(node)).toBeDefined();
    expect(Object.keys(compile(node)).sort()).toEqual(["Compile", "Evaluate", "Pretty", "__internal_object__",]);
  });

  it('evaluate', () => {
    var { node, remainder, err } = parse("@a + @b - @c")
    expect(err).toBeNull();
    expect(remainder).toEqual("");
    var {result, err} = evaluate(compile(node), {"a": 1, "b": 2, "c": 1})
    expect(err).toBeNull();
    expect(result).toEqual(2);
  });

});
