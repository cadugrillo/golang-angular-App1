import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { IMqttMessage } from "ngx-mqtt";
import { MqttClientService } from '../mqttClient.service';

@Component({
  selector: 'app-mqtt-client',
  templateUrl: './mqtt-client.component.html',
  styleUrls: ['./mqtt-client.component.css']
})
export class MqttClientComponent implements OnInit, OnDestroy {

  
  messages: Messages = {message: []};
  message: Message = new Message();
  subscription!: Subscription;
  topic: string = "#"

  constructor(private mqttClientService: MqttClientService) {}

  ngOnInit(): void {
    this.subscribeToTopic();
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();  
}

  subscribeToTopic() {
    this.subscription = this.mqttClientService.topic(this.topic).subscribe((data: IMqttMessage) => {
      //this.message.topic = data.topic;
      //this.message.payload = JSON.parse(data.payload.toString());
      let item = JSON.parse(data.payload.toString());
      //this.message.qos = data.qos;
      //this.messages.message.push(this.message);
      console.log(item);
      
    });
  }
}

class Messages {
  message!: Message[]
}

class Message {
  topic!: string;
  payload!: string;
  qos!: number;
}
