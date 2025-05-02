import sys
import os
from datetime import datetime

# Project information
project = 'Animal Language'
copyright = f'{datetime.now().year}, Animal Language Team'
author = 'Animal Language Team'
version = '1.1.0'  # Feature version
release = '1.1.0'  # Full version string

# General configuration
extensions = [
    'sphinx.ext.autodoc',
    'sphinx.ext.viewcode',
    'sphinx.ext.githubpages',
    'sphinx.ext.intersphinx',
    'sphinx.ext.napoleon',  # Support for Google-style docstrings
    'sphinx_rtd_theme',     # Read the Docs theme
    'myst_parser',          # Support for Markdown
    'sphinx.ext.graphviz',
]

# Add any paths that contain templates here
templates_path = ['_templates']

# The suffix of source filenames
source_suffix = {
    '.rst': 'restructuredtext',
    '.md': 'markdown',
}

# The master toctree document
master_doc = 'index'

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files
exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store']

# The theme to use for HTML and HTML Help pages
html_theme = 'sphinx_rtd_theme'

# Theme options
html_theme_options = {
    'navigation_depth': 4,
    'titles_only': False,
    'logo_only': False,
}

# Add any paths that contain custom static files (such as style sheets)
html_static_path = ['_static']

# Custom sidebar templates
html_sidebars = {
    '**': [
        'relations.html',  # navigation
        'searchbox.html',  # search box
    ]
}

# Output file base name for HTML help builder
htmlhelp_basename = 'AnimalLanguageDoc'