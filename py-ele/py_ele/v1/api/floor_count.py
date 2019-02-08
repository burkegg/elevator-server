# -*- coding: utf-8 -*-
from flask import request, g

from . import Resource
from .. import schemas
from . import elevator

class FloorCount(Resource):

    def get(self):

        return {'count': elevator.floor_count()}, 200, None
