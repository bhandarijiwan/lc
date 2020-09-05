# Copyright (c) Facebook, Inc. and its affiliates.
#
# This source code is licensed under the MIT license found in the
# LICENSE file in the root directory of this source tree.

require "json"

package = "dcllib"
version = "0.0.1"

source = { :git => 'https://github.com/bhandarijiwan/dc/dcl' }
# if version == '1000.0.0'
#   # This is an unpublished version, use the latest commit hash of the react-native repo, which we’re presumably in.
#   source[:commit] = `git rev-parse HEAD`.strip
# else
#   source[:tag] = "v#{version}"
# end

# folly_compiler_flags = '-DFOLLY_NO_CONFIG -DFOLLY_MOBILE=1 -DFOLLY_USE_LIBCPP=1 -Wno-comma -Wno-shorten-64-to-32'
# folly_version = '2020.01.13.00'
# boost_compiler_flags = '-Wno-documentation'

Pod::Spec.new do |s|
  s.name                   = "dcllib"
  s.version                = version
  s.summary                = "-"  # TODO
  s.homepage               = "https://github.com/bhandarijiwan/dc/dcl"
  s.license                = "MIT"
  s.author                 = "Jiwan Bhandari"
  s.platforms              = { :ios => "10.0"}
  s.source                 = source
  s.source_files           = "include/dcllib.h"
  s.preserve_paths         = "lib/dcllib.a"
  s.vendored_frameworks    = "dcllib.xcframewok"
  s.ios.vendored_library   = "lib/dcllib.a"
end
