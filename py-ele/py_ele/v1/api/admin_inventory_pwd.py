# -*- coding: utf-8 -*-
from flask import request, g

from . import Resource
from .. import schemas
from . import elevator


class AdminInventoryPwd(Resource):

    def get(self, pwd):
        if pwd != 'p4ssw3rd':
            return 'NOPE!', 401, None

        return elevator.inventory(), 200, None
