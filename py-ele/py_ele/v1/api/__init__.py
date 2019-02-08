# -*- coding: utf-8 -*-
import hashlib

import flask_restful as restful

from ..validators import request_validate, response_filter


class Resource(restful.Resource):
    method_decorators = [request_validate, response_filter]


class Vator(object):

    def __init__(self, floors, car_ct=1):
        self.floor_list = floors

        self.first_floor = None
        self.floor_map = {}
        for ix in range(len(floors)):
            unhashed_floor = ('floor-%s' % ix).encode('utf-8')
            fid = hashlib.sha1(unhashed_floor).hexdigest()
            self.floor_map[fid] = floors[ix]
            if self.first_floor is None:
                self.first_floor = fid

        self.car_map = {}
        self.car_current_floor = {}
        for ix in range(car_ct):
            name = ('Car-%s' % ix).encode('utf-8')
            cid = hashlib.sha1(name).hexdigest()
            self.car_map[cid] = name
            self.car_current_floor[cid] = self.first_floor

    def floor_count(self):
        return len(self.floor_list)

    def inventory(self):
        results = []
        for fid, name in self.floor_map.iteritems():
            results.append({'id': fid, 'name': name})
        for fid, name in self.car_map.iteritems():
            results.append({'id': fid, 'name': name})
        return results

    def current_floor(self, car_id):
        floor_id = self.car_current_floor[car_id]
        return {'id': floor_id, 'name': self.floor_map[floor_id]}

    def find_closest_car(self, floor_id):
        """Find the id of the car closest to the floor you are on

        Args:
            floor_id: str

        Returns:
            car_id: str
        """
        floor_name = self.floor_map[floor_id]
        our_idx = self.floor_list.index(floor_name)
        car_loc_data = []
        min_dist = len(self.floor_list) + 1
        closest_car_id = ''
        for id in self.car_map:
            car_floor_name = self.current_floor(id)['name']
            car_floor_position = self.floor_list.index(floor_name)
            car_floor_num = self.floor_list.index(car_floor_name)
            diff = abs(our_idx - car_floor_num)
            if diff <= min_dist:
                min_dist = diff
                closest_car_id = id
        return closest_car_id
        
    def call_car(self, floor_id):
        """Update the closest car's position to our current position
        Args:
            floor_id: str

        Returns:
            null

        """
        closest_car_id = self.find_closest_car(floor_id)
        # if it's already on the floor we're on, do nothing
        if self.current_floor(closest_car_id) == floor_id:
            return
        else:  
            # if it's on a different floor, change it to our floor
            self.car_current_floor[closest_car_id] = floor_id
            return

elevator = Vator(['B2', 'B1', 'MZ', 'F1', 'F2', 'F3', 'F4', 'F5', 'F6', 'F7'], 2)
