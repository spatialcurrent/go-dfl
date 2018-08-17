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
evaluate = lib.EvaluateBool
evaluate.argtypes = [c_char_p, c_int, POINTER(POINTER(c_char)), POINTER(c_int)]
evaluate.restype = c_char_p

# Define input and output variables
# Output must be a ctypec_char_p
expression = "(@population > 40) and (@featuretype in [road, highway])"
ctx = {"a": "2", "population": "45", "featuretype": "road"}
output_int_pointer = c_int()

print expression
print ctx

# For explanation see https://mail.python.org/pipermail/python-list/2016-June/709889.html
argv = (POINTER(c_char) * (len(ctx)*2 + 1))()
for i, arg in enumerate([x for t in ctx.iteritems() for x in t]):
    argv[i] =  create_string_buffer(arg.encode('utf-8'))

err = evaluate(expression, len(ctx)*2, argv, byref(output_int_pointer))
if err != None:
    print "error running test.py: %s" % err
    sys.exit(1)

# Output is 0 or 1, so you can easily convert to boolean
output_bool = output_int_pointer.value == 1

# Print output to stdout
print output_bool
