# -*- coding: utf-8 -*-
from flask import request, g

from . import Resource
from .. import schemas


class Welcome(Resource):

    def get(self):

        return {'msg': 'Welcome to the elevator-server example project!'}, 200, None
