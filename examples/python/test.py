# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

from ctypes import *
import sys

# Load Shared Object
# dfl.so must be in the LD_LIBRARY_PATH
# By default, LD_LIBRARY_PATH does not include current directory.
# You can add current directory with LD_LIBRARY_PATH=. python test.py
lib = cdll.LoadLibrary("dfl.so")

# Define Function Definition
fmt = lib.Format
fmt.argtypes = [c_char_p, POINTER(c_char_p)]
fmt.restype = c_char_p

evaluate_bool = lib.EvaluateBool
evaluate_bool.argtypes = [c_char_p, c_int, POINTER(POINTER(c_char)), POINTER(c_int)]
evaluate_bool.restype = c_char_p

evaluate_string = lib.EvaluateString
evaluate_string.argtypes = [c_char_p, c_int, POINTER(POINTER(c_char)), POINTER(c_char_p)]
evaluate_string.restype = c_char_p

# Define input and output variables
# Output must be a ctypec_char_p
expression = "(@population > 40) and (@featuretype in [road, highway])"
ctx = {"a": "2", "population": "45", "featuretype": "road"}
output_int_pointer = c_int()

print expression

output_expression_pointer = c_char_p()

err = fmt(expression, byref(output_expression_pointer))
if err != None:
    print("error: %s" % (str(err, encoding='utf-8')))
    sys.exit(1)

# Convert from ctype to python string
output_expression_value = output_expression_pointer.value

print "Formatted Expression: "+output_expression_value

print ctx

# For explanation see https://mail.python.org/pipermail/python-list/2016-June/709889.html
argv = (POINTER(c_char) * (len(ctx)*2 + 1))()
for i, arg in enumerate([x for t in ctx.iteritems() for x in t]):
    argv[i] =  create_string_buffer(arg.encode('utf-8'))

err = evaluate_bool(expression, len(ctx)*2, argv, byref(output_int_pointer))
if err != None:
    print "error running test.py: %s" % err
    sys.exit(1)

# Output is 0 or 1, so you can easily convert to boolean
output_bool = output_int_pointer.value == 1

# Print output to stdout
print output_bool

output_string_pointer = c_char_p()

err = evaluate_string("concat(@featuretype, ' ', @population)", len(ctx)*2, argv, byref(output_string_pointer))
if err != None:
    print "error running test.py: %s" % err
    sys.exit(1)

output_string_value = output_string_pointer.value

print output_string_value
