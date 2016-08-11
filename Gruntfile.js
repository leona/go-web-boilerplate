var glob = require('glob');
var path = require('path');
var fs = require('fs');
var shell = require('shelljs');
var compressor = require('node-minify');


module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    sass: {
      dist: {
        files: {
          'build/tmp/sass_build.css': 'assets/css/style.scss'
        }
      }
    },
    uglify: {
      options: {
        compress: {
          drop_console: true
        }
      },
      files: {
        files: {
          'build/build.js': ['build/tmp/build_raw.js']
        }
      }
    },
    cssmin: {
      options: {
        shorthandCompacting: false,
        roundingPrecision: -1
      },
      target: {
        files: {
          'build/tmp/minify_build.css': ['build/tmp/sass_build.css']
        }
      }
    },
    imagemin: {                          // Task
      dynamic: {                         // Another target
        files: [{
          expand: true,                  // Enable dynamic expansion
          cwd: 'src/',                   // Src matches are relative to this path
          src: ['**/*.{png,jpg,gif}'],   // Actual patterns to match
          dest: 'dist/'                  // Destination path prefix
        }]
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-imagemin');


  grunt.registerTask('default', ['sass', 'cssmin', 'assets', 'webpack']);

  

  grunt.registerTask('webpack', 'Minifying CSS', function() {
    shell.exec('webpack');
  })
};