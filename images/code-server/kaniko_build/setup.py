from setuptools import setup, find_packages

setup(
    name='kaniko_builder',
    version='1.0.0',
    description='Kaniko pod manager for image builder',
    # url='git@github.com:raizen-analytics-ldt/raizenlib.git',
    author='Raízen Energia Sa',
    author_email='raizen@raizen.ai',
    license='unlicense',
    # packages=find_packages(),
    install_requires=[
          'kubernetes==10.0.1',
      ],
    scripts = [
        'kaniko_build/kaniko-build'
    ],
    zip_safe=False
)