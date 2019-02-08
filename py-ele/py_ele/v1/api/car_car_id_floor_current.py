# -*- coding: utf-8 -*-
from flask import request, g

from . import Resource
from .. import schemas
from . import elevator


class CarCarIdFloorCurrent(Resource):

    def get(self, car_id):
        return elevator.current_floor(car_id), 200, None
