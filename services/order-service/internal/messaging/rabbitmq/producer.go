func (p *Producer) PublishInventoryCheck (order model.Order) error {
	command := InventoryCheckCommand {
		OrderID: order.ID,
		Items : order.Items 
	}

	body, _ := json.Marshal(command)

	return p.channel.Publish(
		"",
		"inventory.check"
		false,
		false,
		amqp.Publishing{
			COntentType:"application/json",
			Body : body,
		},
	)
}