"""Allow running as: python -c "import importlib; importlib.import_module('src.lab.00_get_hands_on.run').main()" """

import importlib

_run = importlib.import_module("src.lab.00_get_hands_on.run")
_run.main()
